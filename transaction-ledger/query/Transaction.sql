-- name: GetTransactionById :one
SELECT 
    id,
    from_account_id,
    amount,
    to_account_id,
    status,
    created_at,
    updated_at
FROM transactions
WHERE id = $1
LIMIT 1;

-- name: UpdateTransactionStatus :exec
UPDATE transactions
SET 
    status = $1,
    updated_at = NOW()
WHERE id = $2;
