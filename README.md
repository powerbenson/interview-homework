### docker build

#### server docker file build
```
docker build -t server -f {YOURE_DIR}/app-server/Dockerfile .
```

#### server client1(socket) file build
```
docker build -t client1 -f {YOURE_DIR}/app-client-socket/Dockerfile .
```

#### server client2(pipe) file build
```
docker build -t client2 -f {YOURE_DIR}/app-client-pipe/Dockerfile .
```

#### server client3(sharedmemory) file build
```
docker build -t client3 -f {YOURE_DIR}/app-client-sharedmemory/Dockerfile .
```

### follow up

1. 錯誤處理很多都只是先捕捉還處理..
2. 連線池可能不要用全域變數比較好(還沒修)
3. 有人斷線時要清除連線池(還沒修)
4. get 或 delete 連線池中的單一連線可能要用鎖
5. sharedmemory 跟 pipe 的機制或許可以不要用這麼多檔案(還沒研究..)
6. pipe 可能不用開關檔案這麼多次(還沒研究)
7. semaphore 的機制跟 syscall 的參數還要再研究..
8. docker 環境沒測過，最好使用本機...
