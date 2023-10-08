#!/bin/bash
mockery --all --dir ../../internal/pkg/useCase/adapters/mock  --keeptree --output ../../internal/pkg/repository/storage/mock --outpkg mockStorage