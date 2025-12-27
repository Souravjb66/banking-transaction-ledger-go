-- name: GetAccountById :one
SELECT 
    id,
    name,
    account_no,
    ifsc_code,
    pin_code,
    balance,
    created_at,
    updated_at,
    deleted_at
FROM accounts
WHERE id = $1
  AND deleted_at IS NULL
LIMIT 1;

-- name: UpdateBalance :exec
UPDATE accounts
SET 
    balance = $1,
    updated_at = NOW()
WHERE id = $2
  AND deleted_at IS NULL;
