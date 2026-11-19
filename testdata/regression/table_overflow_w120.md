| ZK origin | 내 volume 의 origin | 결과 |
|---|---|---|
| 없음 | 없음 | 하위 확인으로. 하위도 모두 비면 첫 commit 에서 origin 을 만든다 |
| 없음 | o | 하위 확인으로. 최신 version 의 timestamp 가 o 와 같으면 ZK 기록을 마저 한다 (3 절) |
| o | o | 하위 확인으로 |
| o | 없음, 또는 p | 복구 필요 |
