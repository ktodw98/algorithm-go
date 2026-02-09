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