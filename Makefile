build:
	hugo --minify

convert:
	go run .

convert-ci:
	go run . --skip-worktree --skip-download

serve:
	hugo serve
