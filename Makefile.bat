@echo on
set input=%1%
if %input% == "" (
    echo "input summary"
) else (
    git add . && git commit -m "%input%" && git push
)

