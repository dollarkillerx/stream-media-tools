# stream-media-tools
流媒体工具  开箱即用


### M3U8下载
``` 
    err := m3u8.M3U8Download("./dow", "https://a-vrv.akamaized.net/evs/4073148fb08c5f2158460ffe74182d52/assets/6bbmnx58kgajfsd_2752571.mp4/index-v1-a1.m3u8?t=st=1567131691~exp=1567218091~acl=/evs/4073148fb08c5f2158460ffe74182d52/assets/6bbmnx58kgajfsd_*~hmac=7a7b46afb1001b6291d2cd66415894a97a8f7507e088be1b8d66db4052800786")
	if err == nil {
		log.Println("下载完毕")
	}
```

### ts 转 mp4
``` 
err := transcoding.TsToMp4("./dow/main.ts", "./out/1.mp4")
```
