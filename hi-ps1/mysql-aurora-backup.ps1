# aurora 项目更新

# $host = '';
# $pswd = '';

$date = Get-Date -Format 'yyMMddHHmmss'
$name = "./aurora-$date.sql"
$user = read-host "Do you need to change the user(root)?"
if ($user -eq ""){
    $user = "root"
}
echo "Will to connect by user: $user."

$database = read-host "Do you wanna change to database(aurora)?"
if ($database -eq ""){
    $database = "aurora"
}

mysqldump -u $user -p --default-character-set=utf8 $database > $name