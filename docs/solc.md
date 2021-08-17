### solc 编译

1. 选择需要版本的solc版本。[https://github.com/ethereum/solidity/](https://github.com/ethereum/solidity/);
2. 官方文档。[https://docs.soliditylang.org/en/v0.8.0/installing-solidity.html](https://docs.soliditylang.org/en/v0.8.0/installing-solidity.html);
3. 注意查看scripts目录下的Dockerfile和其他相关脚本进行编译；
4. cmake时候可能需要`-DUSE_Z3=OFF -DUSE_CVC4=OFF`;