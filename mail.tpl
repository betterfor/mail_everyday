<!DOCTYPE html>
<html>
    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
        <title>一封暖暖的邮件</title>
        <meta name="viewport" content="width=device-width, minimum-scale=1.0, maximum-scale=1.0, user-scalable=no">
        <link rel="stylesheet" href='/stylesheets/style.css' />
    </head>

    <body style="margin:0;padding:0;">
        <div style="width:100%; margin:40px auto; font-size:20px; color: #5f5e5e; text-align:center">
            <span>今日: </span>
            <span style="font-size: 24px; color: #DD4949;">{{ .OneData.Date }}</span>
        </div>
        <div style="width:100%; margin:0 auto; color:#5f5e5e;text-align:center">
            <span style="display: block; color:#676767; font-size:20px">{{ .Tips }}</span>
            <span style="display: block; margin-top: 15px; color: #676767; font-size:15px">近期天气预报</span>
            {{ range .ThreeDays }}
                <div style="display: flex; margin-top:5px; height:30px; line-height:30px; justify-content:space-around;align-items:center;">
                    <span style="width:15%; text-align: center;">{{ .Day }}</span>
                    <div style="width: 25%; text-align: center;">
                        <img style="height:26px; width: 26px; vertical-align: middle;" src='{{ .WeatherImgUrl }}' alt="">
                        <span style="display:inline-block">{{ .WeatherText }}</span>
                    </div>
                    <span style="width:25%; text-align:center">{{ .Temperature }}</span>

                    <div style="width:35%">
                        {{ if eq .PollutionLevel "level_1" }}
                        <span style="display:inline-block; padding:0 8px; line-height:25px; color:#8fc31f; border-radius:15px; text-align:center">{{ .Pollution }}</span>
                        {{ else if eq .PollutionLevel "level_2" }}
                        <span style="display:inline-block; padding:0 8px; line-height:25px; color:#d7af0e; border-radius:15px; text-align:center">{{ .Pollution }}</span>
                        {{ else if eq .PollutionLevel "level_3" }}
                        <span style="display:inline-block; padding:0 8px; line-height:25px; color:#f39800; border-radius:15px; text-align:center">{{ .Pollution }}</span>
                        {{ else if eq .PollutionLevel "level_4" }}
                        <span style="display:inline-block; padding:0 8px; line-height:25px; color:#e2361a; border-radius:15px; text-align:center">{{ .Pollution }}</span>
                        {{ else if eq .PollutionLevel "level_5" }}
                        <span style="display:inline-block; padding:0 8px; line-height:25px; color:#5f52a0; border-radius:15px; text-align:center">{{ .Pollution }}</span>
                        {{ else if eq .PollutionLevel "level_6" }}
                        <span style="display:inline-block; padding:0 8px; line-height:25px; color:#631541; border-radius:15px; text-align:center">{{ .Pollution }}</span>
                        {{- end }}
                    </div>
                </div>
            {{- end }}
        </div>
        <div style="text-align: center; margin: 35px 0;">
            <span style="display:block; margin-top: 55px; color: #676767; font-size:15px">ONE · 一个</span>
            <br/>
            <div style="margin:10px auth; width:85%; color:#5f5e5e">{{ .OneData.Content }}</div>
            <span style="display:block; margin-top: 25px; color: #9d9d9d; font-size:22px">{{ .OneData.Date }}</span>
            <img src="{{ .OneData.ImgUrl }}" style="width: 100%; margin-top: 10px;" alt="">
            <span style="color: #b0b0b0; font-size:13px">{{ .OneData.Type }}</span>
        </div>
    </body>
</html>
