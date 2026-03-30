#!/bin/bash

VERSION="1.0.0"
OUTPUT_DIR="bin"
BINARY_NAME="crm-export"

rm -rf $OUTPUT_DIR
mkdir -p $OUTPUT_DIR

echo "===================================="
echo "编译 1688 CRM 导出工具 v$VERSION"
echo "===================================="

GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o $OUTPUT_DIR/${BINARY_NAME}-darwin-amd64 .
echo "✓ macOS amd64 编译完成"

GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o $OUTPUT_DIR/${BINARY_NAME}-darwin-arm64 .
echo "✓ macOS arm64 编译完成"

GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o $OUTPUT_DIR/${BINARY_NAME}-linux-amd64 .
echo "✓ Linux amd64 编译完成"

GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o $OUTPUT_DIR/${BINARY_NAME}-linux-arm64 .
echo "✓ Linux arm64 编译完成"

GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o $OUTPUT_DIR/${BINARY_NAME}-windows-amd64.exe .
echo "✓ Windows amd64 编译完成"

echo ""
echo "===================================="
echo "编译完成！输出文件："
echo "===================================="
ls -lh $OUTPUT_DIR/

echo ""
echo "文件大小对比："
du -sh $OUTPUT_DIR/*
