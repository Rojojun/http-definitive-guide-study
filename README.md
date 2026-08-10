# HTTP - The Definitive Guide Study

`HTTP - The Definitive Guide`를 하루 최소 60분씩 학습하면서 ChatGPT와 Codex가 동일한 진도와 요약을 이어서 사용하기 위한 비공개 저장소입니다.

## 저장소 구성

- `learning-ledger.md`: 학습 계획, 현재 진도, 교정한 오해, 복습 큐, 단원별 Summary
- `Http - The Definitive Guide.pdf`: 개인 학습용 원서 PDF
- `plugins/study-coach-sync/`: ChatGPT와 Codex 공용 학습 코치 플러그인
- `.agents/plugins/marketplace.json`: 저장소 기반 플러그인 마켓플레이스

원서 PDF를 포함하므로 이 저장소는 비공개 상태를 유지합니다.

## ChatGPT Work 웹에서 가져오기

먼저 ChatGPT Work의 새 대화에서 도구 메뉴에 `GitHub Study Ledger` 연결을 추가합니다. 그다음 아래 요청을 한 번 실행합니다.

```text
@plugin-creator

Import the existing plugin from my private GitHub repository
Rojojun/http-definitive-guide-study at plugins/study-coach-sync.

Use my registered MCP connection:
plugin_asdk_app_6a797ae5ad908191806720b35d73c2d4

Preserve the bundled diagnose-and-teach skill and all reference files.
Create or update my Personal marketplace entry, validate the plugin,
and make it installable in ChatGPT Work.
```

생성이 끝나면 `Work → Plugins → Personal → Created by me`에서 `Study Coach Sync`를 설치하고 새 대화를 시작합니다.

```text
@diagnose-and-teach

GitHub 학습 원장을 읽고 현재 진도와 복습 큐를 확인한 뒤 다음 학습을 시작해줘.
학습이 끝나면 진도와 Summary를 learning-ledger.md에 커밋해줘.
```

## Codex에서 저장소 마켓플레이스 사용하기

```text
codex plugin marketplace add Rojojun/http-definitive-guide-study --ref main
codex plugin add study-coach-sync@http-study
```

플러그인을 설치하거나 갱신한 뒤에는 새 세션에서 시작합니다.

## 동기화 원칙

- 원격 `main`의 `learning-ledger.md`를 기준으로 시작합니다.
- 쓰기 직전에 최신 파일 SHA를 다시 확인합니다.
- 현재 세션의 변경분만 병합하고 알 수 없는 기존 내용은 보존합니다.
- 강제 푸시하지 않습니다.
- 토큰, 비밀번호 및 학습과 무관한 파일은 업로드하지 않습니다.
