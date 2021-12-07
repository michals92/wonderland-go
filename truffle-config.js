module.exports = {
  plugins: [
      "@chainsafe/truffle-plugin-abigen"
  ],
  compilers: {
      solc: {
          version: "0.8.10",    // Fetch exact version from solc-bin (default: truffle's version)
      }
  }
}