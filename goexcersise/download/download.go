func (d *ResDownloader) Start() {
	go func() {
		idleTimes := 0
		needOss := d.client.Config.GetBool("download.oss.enabled")
		ossConf := d.client.Config.GetString("download.oss.config")
		retryTimeLimit := d.client.Config.GetInt("download.retry.time_limit")
		ossRetryTimeLimit := d.client.Config.GetInt("download.oss.retry_time_limit")
		downloadProxy := d.client.Config.GetString("download.proxy")
		downloadConcurrentLimit := d.client.Config.GetInt("download.concurrent_limit")
		downloadDelayLimit := d.client.Config.GetInt("download.delay_limit")
		clientId := d.client.ID
		// 用通道限制同时最多N个下载任务
		limiter := make(chan int, downloadConcurrentLimit)
		for {
			select {
			case job := <-d.jobCh:
				// 申请一个下载通道
				limiter <- 1
				idleTimes = 0
				go func(job CrawlerCommon.Job) {
					// 返回时释放一个下载通道
					defer func() {
						<-limiter
					}()
					// 随机休息一段时间，避免同时过多下载任务
					time.Sleep(time.Second * time.Duration(rand.Intn(downloadDelayLimit)))
					targetUrl := job.QSA["url"]
					proxy := "http://127.0.0.1:108"//job.QSA["proxy"]
					referer := job.Referer
					var task CrawlerCommon.Task
					json.Unmarshal([]byte(job.QSA["task"]), &task)
					if task.Site == "" {
						log.Errorf("skip download task without site %#v", task)
						//panic("where is this task from")
					}
					if proxy == "" && downloadProxy != "" {
						proxy = downloadProxy
					}

					log.Debugf("download with proxy %#v", proxy)
					dd := Common.NewDownloader()
					dd.SetNeedOSS(needOss).WithOssConf(ossConf).WithRetryTimeLimit(retryTimeLimit).WithOssRetryTimeLimit(ossRetryTimeLimit).WithProxy(proxy)
					dd.Download(d.downloadTo, targetUrl, referer, func(e error, remoteUrl string, localFilename string, cdnUrl string, hash string, bodyHash string) {
						now := time.Now()
						if e != nil {
							downloaderStats.Add("fail", 1)
							log.Errorf("download fail for  %s for %s", remoteUrl, e)
							d.client.writeCh <- map[string]string{
								"topic":     "medusa_images",
								"id":        strings.Join([]string{task.ID, "batch"}, ":::"),
								"batchMode": "yes",
								"payload": Common.JsonEncode(map[string]interface{}{
									"status":       "fail",
									"status_extra": fmt.Sprintf("download fail for %s at %s %s", e, clientId, now.Format(time.RFC3339)),
									"task":         task,
								}),
							}
							return
						}

						downloaderStats.Add("success", 1)
						log.Infof("downloaded %s local:%s cdn: %s", remoteUrl, localFilename, cdnUrl)
						if d.client != nil {

							d.client.writeCh <- map[string]string{
								"topic":     "medusa_images",
								"id":        strings.Join([]string{task.ID, "batch"}, ":::"),
								"batchMode": "yes",
								"payload": Common.JsonEncode(map[string]interface{}{
									"item": Common.MergeMapSS(job.ENV, map[string]string{
										"url":           remoteUrl,
										"cdn_url":       cdnUrl,
										"url_hash":      hash,
										"body_hash":     bodyHash,
										"download_time": now.Format(time.RFC3339),
										"downloaded_at": strconv.FormatInt(now.Unix(), 10),
										"downloaded_by": clientId,
									}),
									"status":       "success",
									"status_extra": "done",
									"task":         task,
								}),
							}
						}
					})
				}(job)

			case sig := <-d.sigCh:
				log.Debugf("quit downloader for %d", sig)
			default:
				idleTimes++
				if idleTimes >= d.maxIdleTimes {
					idleTimes = 0
				}
				time.Sleep(time.Second * time.Duration(math.Pow(2, float64(idleTimes))))
			}
		}
	}()
}