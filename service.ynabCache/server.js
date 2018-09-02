const ynabApi = require('ynab');
const http = require('http');
const port = 8080;
var cache = null


async function getData(api_token, budget_id, cache_days) {
  // Query YNAB for data and write to cache.
  
  var ynab = new ynabApi.api(api_token);
 
  var dateSince = new Date();
  dateSince.setDate( new Date().getDate() - cache_days) ;
  transactions = await ynab.transactions.getTransactions(budget_id, sinceDate=dateSince)


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
  // logging routine.

  let t = (new Date).getTime();
  console.log(`${t}:\t${content}`);
}

function main() {

  let api_token = process.env.ynab_api_token; 
  let budget_id = process.env.ynab_budget_id
  let TTL = process.env.cache_ttl;
  let cache_days = process.env.cache_days
  let LastCache = 0; 

  console.log("Cache Service Started")
  console.log("Budget id: "+ budget_id)
  console.log("TTL: " + TTL )


  // create HTTP server for cache
  http.createServer(function (req, res) {

    // GET request
    if (req.method == 'GET') { 
      log("Cache Request.");
      // Check for empty or expired cache. 
      if (cache == null || (new Date).getTime() 
        > (LastCache + TTL))
      {
        log("Cache Miss.");
        getData(api_token, budget_id, cache_days);
        LastCache = (new Date).getTime();
        res.writeHead(204, {'Content-Type': 'application/json'});
        res.end();
        return
      }

      res.writeHead(200, {'Content-Type': 'application/json'});
      res.end(cache);
      log("Cache Served.");
    }
    // POST request
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
