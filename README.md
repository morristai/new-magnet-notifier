# Rarbg Notifier

- https://github.com/robfig/cron
- https://github.com/DudeofA/kyBot/blob/master/kylixor.go
- https://github.com/go-redis/redis
- 

## TODO

1. imdb script call
3. check 4K
4. parse movie info

## Parser (splite every dot after year! DPP5.1?)

### Must
1. Title: `(^\S*)`
2. Year: `((?:(?:19|20)[0-9]{2}))`  1900 to 2099

### Optional

- SUBBE
- HDR
- ATVP (apple TV)

- Resolution: `((?:720|1080|2160)p)`

- Source: AMZN, NF, DIS?, NOW (nowtv), 
- Formats: WEBRip, WEB, WEBRip, BluRay, WEB-DL, 
- Audio: DDP5.1, TrueHD.7.1, Atmos-*
- Encoding: x264-TEPES, H264-NAISU, x265, x264-ION10, x265-RARBG, x264-RARBG, H265-SLOT, x264-NOGRP, x264-CM
- Language: KOREAN

example
- BluRay.REMUX.HEVC.DTS-HD.MA.TrueHD.7.1.Atmos-FGT
- Alive.2020.KOREAN.BRRip.x264-VXT (language)
- Endgame.1983.ITALIAN.BRRip.x264-VXT (language)
- Volumes.of.Blood.Horror.Stories.2016.BRRip.x264-ION10 (no resolution)
- The.Unforgivable.2021.1080p.NF.WEBRip.DDP5.1.Atmos.x264-CM (Netflix)