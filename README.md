# 🐹 Kai's Go Algorithm Journey

Go 언어 숙달과 효율적인 문제 해결 능력을 기르기 위한 알고리즘 저장소입니다.  
단순한 풀이를 넘어 **Idiomatic Go**와 **성능 최적화**를 지향하며, 테스트 주도 개발(TDD) 방식으로 학습합니다.

<img src="./assets/profile.png" width="300"> <img src="./assets/golang.png" width="300"> 

---
## 🛠️ Tech Stack & Principles

- **Language:** Go (1.20+)
- **Strategy:**
  - **Table-driven Testing:** 모든 풀이는 Go의 표준 테스트 라이브러리로 검증합니다.
  - **Zero-Dependency:** 외부 라이브러리 없이 Go 표준 라이브러리만 사용합니다.

---
## 🚀 Workflow (How to use)

`Makefile`을 통해 반복적인 작업을 자동화하여 문제 해결에만 집중합니다.

### 1. 문제 생성 (New Problem)
플랫폼별 명령어를 사용하여 표준화된 폴더 구조를 생성합니다.
```bash
# LeetCode: 번호(n)와 이름(name) 지정
make new-lc n=1 name=two_sum

# Baekjoon: 번호(n)와 이름(name) 지정
make new-boj n=1000 name=plus

# Programmers: 이름(name) 지정
make new-pg name=target_number
```

### 2. 특정 문제 테스트 및 실행 (Test & Run)
불필요한 전체 테스트를 피하고, 현재 작업 중인 문제만 타겟팅합니다.

```bash
# 릿코드: 문제 번호로 테스트 코드 실행
make test-lc n=1

# 백준: 문제 번호로 코드 직접 실행 (표준 입출력 확인)
make run-boj n=1000

# 경로 지정: 특정 경로의 패키지만 테스트
make test p=leetcode/p0001_two_sum
```

### 3. Project Structure
```bash
.
algorithm-go/
├── internal/             # (공통) 내가 만든 라이브러리 (직접 구현한 자료구조 등)
│   └── ds/               # Stack, Queue, Heap, Set 등
├── leetcode/             # 릿코드: 함수만 제출하는 방식
│   └── p0001_two_sum/
├── boj/                  # 백준: main 패키지와 입출력이 필요한 방식
│   └── b1000_a_plus_b/
│       └── main.go       # 백준은 파일 하나에 모든 코드를 넣어야 함
├── programmers/          # 프로그래머스: 솔루션 함수 중심
│   └── level1_test/
├── Makefile              # 플랫폼별 새 문제 생성 스크립트
└── go.mod
```

### 4. Commit Convention

#### 1. 커밋 메시지 구조
Conventional Commits 형식을 알고리즘 학습에 맞춰 변형하여 사용합니다.

```
<type>(<scope>): <subject> [#issue]

<body> (선택 사항)
```

- **type**: 커밋의 성격
- **scope**: 플랫폼 또는 문제 번호/이름
- **subject**: 변경 사항에 대한 요약 (한글 또는 영어)

#### 2. 권장하는 Type 및 Scope

**Type 정의**

| Type | 설명 |
|------|------|
| `solve` | 새로운 문제를 처음으로 해결했을 때 (기능 구현에 해당) |
| `refactor` | 로직을 깔끔하게 개선하거나 Idiomatic Go 스타일로 수정했을 때 |
| `perf` | 벤치마크 기반 성능 최적화 (메모리 할당 감소, 시간 복잡도 개선 등) |
| `docs` | README 수정, 문제 설명(README) 추가 등 문서 관련 변경 |
| `test` | 테스트 케이스 추가 또는 테스트 코드 수정 |
| `internal` | internal/ds 등 공통 자료구조 라이브러리 개발/수정 |
| `chore` | Makefile 수정, 패키지 매니저(go.mod) 설정 변경 등 |

**Scope 정의**

- `lc`: LeetCode
- `boj`: Baekjoon
- `pg`: Programmers
- `ds`: Data Structures (internal 패키지)

#### 3. 커밋 예시

**새로운 문제 해결**
```
solve(lc): p0001 two sum
solve(boj): b1000 a+b 입출력 템플릿 적용
```

**성능 최적화 및 리팩토링**
```
perf(lc): p0001 map을 활용하여 O(n^2)에서 O(n)으로 최적화
refactor(ds): generic stack에 포인터 리시버 적용
```

**문서 및 테스트**
```
docs: 전체 학습 현황판 업데이트
test(pg): 타겟 넘버 테스트 케이스 3종 추가
docs: update problem solved stats [skip ci]
```

#### 4. 원칙 (Principles)

- **원자적 커밋 (Atomic Commits)**: 한 커밋에는 하나의 문제 해결만 포함합니다. 여러 문제를 풀었더라도 각각 따로 커밋하세요.
- **이유가 담긴 커밋**: 특히 `perf`나 `refactor`의 경우, 무엇을 바꿨는지보다 **왜(Why)** 바꿨는지를 제목이나 본문에 짧게 적어주는 것이 좋습니다.
  예: `"슬라이스 재할당 방지를 위해 cap 지정"`
- **자동화 커밋 처리**: GitHub Actions에 의해 발생하는 커밋은 메시지에 `[skip ci]`를 포함하여 불필요한 워크플로우 실행을 방지합니다.