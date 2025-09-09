# generate-grpc.ps1
protoc `
  --go_out=. `
  --go-grpc_out=. `
  --go_opt=paths=source_relative `
  --go-grpc_opt=paths=source_relative `
  internal/grpc/wallet.proto

Write-Host "✅ gRPC code generated successfully!" -ForegroundColor Green