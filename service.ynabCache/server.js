const ynabApi = require('ynab');
const http = require('http');
const port = 8080;
var cache = null


async function getData(api_token, budget_id) {
  // Query YNAB for data  
  
  var ynab = new ynabApi.api(api_token);

  transactions = await ynab.transactions.getTransactions(budget_id);

  var options = {
    hostname: 'localhost',
    port: 8080,
    path: '/',
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    }
  }

  var req = http.request(options, function(r){
    r.on('end',function(){
      log('Cache Sent.');
    }) 
    
  });

  req.end(JSON.stringify(transactions));



}

function log(content){

  let t = (new Date).getTime();

  console.log(`${t}:\t${content}`);
}

function main() {

  let api_token = process.env.ynab_api_token; 
  let budget_id = process.env.ynab_budget_id
  let TTL = process.env.cache_ttl;
  let LastCache = 0; 

  http.createServer(function (req, res) {

    if (req.method == 'GET') { 
      log("Cache Request.");
      
      if (cache == null || (LastCache + TTL) 
        > (new Date).getTime()) {
        
        log("Cache Miss.");
        getData(api_token, budget_id);
        LastCache = (new Date).getTime();


        res.writeHead(204, {'Content-Type': 'application/json'});

        res.end();
        return
      }

      res.writeHead(200, {'Content-Type': 'application/json'});

      res.end(cache);
      
      log("Cache Served.");
    }

    else if (req.method == 'POST') {
      log("Cache Update Started.");

      cache = "";
      req.on('readable', function() {

        chunk = req.read();

        if (chunk != null) {
          cache += chunk;
        }
      });

      log("Cache Updated Finished.");
    }
  }).listen(port);
}



main()

