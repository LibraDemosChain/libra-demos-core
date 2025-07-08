// Example JS API usage for LibraDemosChain
fetch('http://localhost:1317/cosmos/gov/v1beta1/proposals')
  .then(res => res.json())
  .then(data => console.log(data));
