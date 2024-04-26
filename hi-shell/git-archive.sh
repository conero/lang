# 使用 git archive 打包当前仓库代码
# 命名格式 `$project-$hash-$latestdata--$branch.tar.gz`

# 读取项目名称
prjname=$(basename "$PWD")

# 当前仓库的 hash 值
curhash=$(git rev-parse HEAD)
curhash=${curhash:0:10}

# 分支名称
branch=$(git rev-parse --abbrev-ref HEAD)
branch=${branch//\//-}

last_date=$(git log -n 1 --format="%cd" --date=format:'%Y%m%d')

# 目录检查
if [ ! -d dist ]; then
  mkdir dist
  # 忽略任何东西
  echo "*" > dist/.gitignore
fi

name="dist/$prjname-$curhash-$last_date--$branch.tar.gz"

git archive --format=tar.gz -o $name HEAD

echo -e " 仓库代码已打包\n  \033[32m$name\033[0m"