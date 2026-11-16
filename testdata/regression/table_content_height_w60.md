| 방식 | 예 | 결과 |
|------|-----|------|
| **소스 재컴파일** (`$<TARGET_PROPERTY:X,SOURCES>`) | `ircomm_test`, `srch_is_test`, `srch_datapath_test`, `srch_rerank_test`, `srch_infosrch_test` | hidden preset이 적용되지 않아 `_hidden` 계열 심볼까지 전부 export된다 |
| **OBJECT 직접 링크** (오브젝트 내장) | `srch_server_test`←`srch_server_export`, `srch_rerank_test`←`ircomm_export`/`ircomm_hidden` | 해당 모듈 코드가 그 `.so` 안에 사본으로 존재 |
