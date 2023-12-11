-- name: CreateSignature :one
INSERT INTO signatures (
    user_id,
    signature,
    answers,
    questions,
    timestamp
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
) RETURNING *;

-- name: GetSignatureByUserIdAndSignature :one
SELECT * FROM signatures 
WHERE user_id = $1 AND signature = $2;