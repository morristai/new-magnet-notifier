# New Magnet Notifier
> **Disclaimer**: This repository is for educational purposes only. This tool do not promote or encourages any illegal pirated action and it was not made for criminal purposes.

A Discord bot that auto sends notifications when there's new movies

<img src="https://user-images.githubusercontent.com/10728152/154961582-20c65070-7010-4f61-ba77-f5622a64733d.png" width="400" />

## Components
### Parser
#### Required
1. Title: `(^\S*)`
2. Year: `((?:(?:19|20)[0-9]{2}))`  1900 to 2099
#### Optional
- SUBBE
- HDR
- ATVP (apple TV)
- Resolution: `((?:720|1080|2160)p)`
- Source: AMZN, NF, DIS?, NOW (nowtv), 
- Formats: WEBRip, WEB, WEBRip, BluRay, WEB-DL, 
- Audio: DDP5.1, TrueHD.7.1, Atmos-*
- Encoding: x264-TEPES, H264-NAISU, x265, x264-ION10, x265-RARBG, x264-RARBG, H265-SLOT, x264-NOGRP, x264-CM
- Language: KOREAN, SPANISH

#### example
- BluRay.REMUX.HEVC.DTS-HD.MA.TrueHD.7.1.Atmos-FGT
- Alive.2020.KOREAN.BRRip.x264-VXT (language)
- Endgame.1983.ITALIAN.BRRip.x264-VXT (language)
- Volumes.of.Blood.Horror.Stories.2016.BRRip.x264-ION10 (no resolution)
- The.Unforgivable.2021.1080p.NF.WEBRip.DDP5.1.Atmos.x264-CM (Netflix)

### TODO
1. Better error handling
2. Add Ansible
3. Add Github Action
4. Check 4K
5. Genre emoji map
6. Check all resolutions of the film

## Others problems during development
### Why not use Ordered Map to keep tracking notified list?
As an issue pointed out here: [GitHub Issue](https://github.com/elliotchance/orderedmap/issues/12)
1. The order will be lost when serializing (and unserializing). There's no way around that because **Go sorts maps for JSON output**.
2. All of your keys must be strings.

## Workflow
1. main -> run bot main process discord/bot.go -> trigger cron job (task/task.go)
2. task(read config) -> processor (request Rarbg) -> parser/parser.go
