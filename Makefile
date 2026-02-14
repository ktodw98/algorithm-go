# Makefile
.PHONY: help new-lc new-boj new-pg test test-lc test-pg run-boj

# 기본 도움말
help:
	@echo "🚀 문제 생성 명령어:"
	@echo "  make new-lc n=번호 name=이름   : LeetCode 생성"
	@echo "  make new-boj n=번호 name=이름  : 백준 생성"
	@echo "  make new-pg name=이름         : 프로그래머스 생성"
	@echo ""
	@echo "🧪 테스트 및 실행 명령어:"
	@echo "  make test p=경로              : 특정 폴더 테스트 (예: make test p=leetcode/p0001_two_sum)"
	@echo "  make test-lc n=번호           : 릿코드 번호로 테스트 (예: make test-lc n=1)"
	@echo "  make run-boj n=번호           : 백준 번호로 바로 실행 (예: make run-boj n=1000)"

# --- 문제 생성 로직 ---

new-lc:
	@$(eval DIR := leetcode/p$(shell printf "%04d" $(n))_$(name))
	@mkdir -p $(DIR)
	@echo "package p$(shell printf "%04d" $(n))\n\nfunc solve() {\n\n}" > $(DIR)/solution.go
	@echo "package p$(shell printf "%04d" $(n))\n\nimport \"testing\"\n\nfunc TestSolve(t *testing.T) {\n\n}" > $(DIR)/solution_test.go
	@echo "package p$(shell printf "%04d" $(n))\n\nimport \"testing\"\n\nfunc BenchmarkSolve(b *testing.B) {\n\tfor i := 0; i < b.N; i++ {\n\t\tsolve()\n\t}\n}" > $(DIR)/solution_benchmark_test.go
	@echo "# LeetCode $(n). $(name)\n\n- [링크](https://leetcode.com/problems/$(subst _,-,$(name))/)" > $(DIR)/README.md
	@echo "✅ LeetCode $(n)번 생성 완료"

new-boj:
	@$(eval DIR := boj/b$(n)_$(name))
	@mkdir -p $(DIR)
	@echo "package main\n\nimport (\n\t\"bufio\"\n\t\"os\"\n)\n\nfunc main() {\n\treader := bufio.NewReader(os.Stdin)\n\twriter := bufio.NewWriter(os.Stdout)\n\tdefer writer.Flush()\n}" > $(DIR)/main.go
	@echo "# 백준 $(n). $(name)\n\n- [링크](https://www.acmicpc.net/problem/$(n))" > $(DIR)/README.md
	@echo "✅ 백준 $(n)번 생성 완료"

new-pg:
	@$(eval DIR := programmers/$(name))
	@mkdir -p $(DIR)
	@echo "package programmers\n\nfunc solution() {\n\n}" > $(DIR)/solution.go
	@echo "package programmers\n\nimport \"testing\"\n\nfunc TestSolution(t *testing.T) {\n\n}" > $(DIR)/solution_test.go
	@echo "✅ 프로그래머스 생성 완료"

# --- 테스트 및 실행 로직 ---

# 특정 경로 테스트
test:
	@if [ -z "$(p)" ]; then echo "❌ 경로를 입력하세요 (p=path)"; else go test -v ./$(p); fi

# 릿코드 번호로 테스트
test-lc:
	@$(eval TARGET := $(shell find leetcode -name "p$(shell printf "%04d" $(n))*" -type d))
	@if [ -z "$(TARGET)" ]; then echo "❌ 문제를 찾을 수 없습니다."; else go test -v ./$(TARGET); fi

# 백준 번호로 실행 (입력은 터미널에 직접 입력하거나 < input.txt 활용)
run-boj:
	@$(eval TARGET := $(shell find boj -name "b$(n)*" -type d))
	@if [ -z "$(TARGET)" ]; then echo "❌ 문제를 찾을 수 없습니다."; else go run $(TARGET)/main.go; fi