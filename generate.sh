#!/bin/bash


# Prompt the user for their name
echo "Please enter your name:"
read name

#table_name construction
table_name=$(echo "$name" |  cut -c 3-)
table_name=$(echo "$table_name"'s')

#model_name construction
model_name=$(echo "$name" | awk -F'_' '{for(i=1;i<=NF;i++) $i=toupper(substr($i,1,1))substr($i,2);}1' | tr -d ' ' | cut -c 2-)

#logger_name construction
logger_name=$(echo "$name" | awk -F'_' '{for(i=1;i<=NF;i++) $i=substr($i,1,1)substr($i,2);}1' | tr -d '_' | cut -c 3-)

#variable_name
variable_name=$(echo "$model_name" | sed 's/\b\(.\)/\L\1/g')


cp models/a_user_group.go "models/$name.go"

#sleep 2

sed -i '' "s/UserGroup/$model_name/g" "models/$name.go"
sed -i '' "s/user_groups/$table_name/g" "models/$name.go"
sed -i '' "s/user group/$logger_name/g" "models/$name.go"
sed -i '' "s/userGroup/$variable_name/g" "models/$name.go"


#cp routes/auth.go "routes/$name.go"

#cp controllers/auth.go "controllers/$name.go"

echo "Hello, $name!"