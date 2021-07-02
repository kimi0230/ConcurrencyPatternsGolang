# fan-in fan-out

## fan-in
多對一, 類似 leaky bucket. 速度取決於消費者(擋 request)

## fan-out
一對多, 類似 token bucket. 速度取決於生產者(給 redis/db 等後端資源)