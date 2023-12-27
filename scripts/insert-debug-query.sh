#!/bin/bash

# PostgreSQL接続設定
PGUSER="postgres"
PGPASSWORD="password"
PGHOST="localhost"
PGPORT="${POSTGRES_PORT:-5432}"
PGDATABASE="prel"

export PGUSER PGPASSWORD PGHOST PGPORT PGDATABASE

generate_random_string() {
    cat /dev/urandom | LC_ALL=C tr -dc 'a-zA-Z0-9' | fold -w ${1:-32} | head -n 1
}

generate_random_date() {
    if [[ "$OSTYPE" == "darwin"* ]]; then
        gdate -d "2023-12-01 + $((RANDOM % 365)) days + $((RANDOM % 86400)) seconds" --iso-8601=seconds
    else
        date -d "2023-12-01 + $((RANDOM % 365)) days + $((RANDOM % 86400)) seconds" --iso-8601=seconds
    fi
}

num=${1:-100}

declare -a user_ids
for i in $(seq 1 $((num * 2))); do
    user_ids[$i]="$(generate_random_string 10)"
done

# リクエストデータの生成
request_values=""
for i in $(seq 1 $num); do
    id=$(generate_random_string 10)
    requester_user_id=${user_ids[$((RANDOM % (num * 2) + 1))]}

    statuses=("approved" "pending" "rejected")
    status=${statuses[$RANDOM % ${#statuses[@]}]}
    if [ "$status" == "pending" ]; then
        judger_user_id="NULL"
        judged_at="NULL"
    else
        judger_user_id="'${user_ids[$((RANDOM % (num * 2) + 1))]}'"
        while [ "$requester_user_id" == "$judger_user_id" ]; do
            judger_user_id="'${user_ids[$((RANDOM % (num * 2) + 1))]}'"
        done
        judged_at="'$(generate_random_date)'"
    fi

    project_id=$(generate_random_string 10)
    iam_roles="role_$i"
    reason="reason_$i"
    requested_at="'$(generate_random_date)'"
    expired_at="'$(generate_random_date)'"

    request_values="$request_values ('$id', '$requester_user_id', $judger_user_id, '$status', '$project_id', '$iam_roles', '$reason', $requested_at, $expired_at, $judged_at),"
done

# ユーザーデータの生成
user_values=""
for i in $(seq 1 $((num * 2))); do
    id=${user_ids[$i]}
    google_id=$(generate_random_string 10)
    email=${user_ids[$i]}
    is_available="true"
    role="requester"
    session_id=$(generate_random_string 10)
    session_expired_at="'$(generate_random_date)'"
    last_signin_at="'$(generate_random_date)'"

    user_values="$user_values ('$id', '$google_id', '$email', $is_available, '$role', '$session_id', $session_expired_at, $last_signin_at),"
done

# 末尾のカンマを削除
request_values=${request_values%?}
user_values=${user_values%?}

# バルクインサートの実行
echo "INSERT INTO users (id, google_id, email, is_available, role, session_id, session_expired_at, last_signin_at) VALUES $user_values;" | psql
echo "INSERT INTO requests (id, requester_user_id, judger_user_id, status, project_id, iam_roles, reason, requested_at, expired_at, judged_at) VALUES $request_values;" | psql
