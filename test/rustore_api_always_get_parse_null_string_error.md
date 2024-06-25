## RuStore API always get "parse null string" error

### Go SDK

repo https://gitflic.ru/project/rustore/rustoreapi

The official Go SDK uses keyId but not companyId, but we want to use companyId,
so I did a little hack in source code.

```go
// client.go
signatureValue := base64.StdEncoding.EncodeToString(signatureBytes)

resultMap := map[string]string{
    // official SDK is keyId, I changed it to companyId
    "companyId": keyID,
    "signature": signatureValue,
    "timestamp": timestamp,
}

ResultJSON, err := json.Marshal(resultMap)

return string(ResultJSON), err
```

Here below is my test code

```go
// rustoreapi_test.go
const (
	keyID         = "2318908607" // real company id
	companyID     = "2318908607" // real company id
	privateKey    = "..." // fake private key
	packageName   = "com.raven.tank.rustore" // real package name
	purchaseToken = "1000005247.2520743208" // real purchase token
)

func TestRuStoreAPI(t *testing.T) 
    // keyId & companyId are the same
	c := client.New(keyID, privateKey, companyID)
	err := c.Auth()
	if err != nil {
		t.Error(err)
		return
	}

	p := payments.New(c, packageName)
	ctx, cancel := context.WithTimeout(context.Background(), client.TimeOutSeconds*10*time.Second)
	defer cancel()
	info, err := p.GetPaymentInfo(ctx, purchaseToken)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(info)
}
```

But the result is always

```
/usr/local/opt/go/libexec/bin/go tool test2json -t /Users/caiknife/Library/Caches/JetBrains/GoLand2023.3/tmp/GoLand/___TestRuStoreAPI_in_github_com_caiknife_mp3lister_test.test -test.v -test.paniconexit0 -test.run ^\QTestRuStoreAPI\E$
=== RUN   TestRuStoreAPI
    rustoreapi_test.go:36: {"code":"ERROR","message":"Cannot parse null string","body":{"invoice_id":"","invoice_date":"","invoice_status":"","application_code":"","application_name":"","owner_code":"","owner_name":"","payment_info":{"payment_date":"","payment_id":"","payment_params":null,"loyalty_info":null,"card_id":"","paysys_code":"","masked_pan":"","expiry_date":"","payment_system":"","payment_system_image":"","paysys_image":null,"payment_way":"","payment_way_code":"","payment_way_logo":"","bank_info":{"bank_name":"","bank_country_code":"","bank_country_name":"","bank_image":""},"device_info":null,"name":"","cardholder":"","image":"","paysys":""},"payment_methods":null,"error":{"user_message":"","error_description":"","error_code":0},"invoice":{"delivery_info":{"delivery_type":null,"address":{"country":null,"city":null,"address":null},"description":null},"invoice_params":null,"purchaser":{"email":"","phone":null,"contact":null},"order":{"order_id":"","order_number":"","order_date":"","service_id":"","expiration_date":"","tax_system":0,"trade_name":null,"visual_name":"","org_name":"","org_inn":"","visual_amount":"","order_bundle":null,"amount":0,"currency":"","purpose":"","description":"","language":""}},"image":""},"timestamp":"2024-04-15T05:20:59.952468469+03:00"}
--- PASS: TestRuStoreAPI (2.89s)
PASS
```

Did I use wrong parameters?

I also found a Node.js sdk on GitHub, https://github.com/piavart/rustore-client, the result is exactly the same as above.

``` js
// source code rustore.js
const {RuStoreClient, RS_InvoiceStatus} = require("@piavart/rustore-client")
const privateKey = "MIIEvAIB..."
const companyID = 23
const client = new RuStoreClient(privateKey, companyID)

class RuStore {
    async getPurchase(purchaseToken) {
        try {
            let purchase = await client.getPurchase(purchaseToken)
            return purchase
        } catch (e) {
            throw e
        }
    }

    async getSubscription(purchaseToken) {
        try {
            let purchase = await client.getSubscription(purchaseToken)
            return purchase
        } catch (e) {
            throw e
        }
    }

    checkStatus(status) {
        try {
            return status === RS_InvoiceStatus.Confirmed
        } catch (e) {
            throw e
        }
    }
}

const rustore = new RuStore()

module.exports = {
    rustore
}
```

```js
//test code rustore.test.js
const {test, describe, expect} = require("@jest/globals")
const {rustore} = require("./rustore");

describe("rustore", () => {
    test("getPurchase", async () => {
        try {
            // this purchaseToken is a real token in my app
            let purchaseToken = "1000005247.2520743208"
            let result = await rustore.getPurchase(purchaseToken)
            console.log(result)
        } catch (e) {
            console.error(e.stack)
            console.error(e.data)
        }
    })
})
```

```js
console.error
    Error: Cannot parse null string
        at RuStoreClient.request (/Users/caiknife/WebstormProjects/tank_server/node_modules/@piavart/rustore-client/dist/client.js:45:19)
        at process._tickCallback (internal/process/next_tick.js:68:7)

      11 |             console.log(result)
      12 |         } catch (e) {
    > 13 |             console.error(e.stack)
         |                     ^
      14 |             console.error(e.data)
      15 |         }
      16 |     })

      at Object.test (pub/utils/pay/rustore.test.js:13:21)

  console.error
    { code: 'ERROR',
      message: 'Cannot parse null string',
      body: null,
      timestamp: '2024-04-12T14:15:19.021183821+03:00' }

      12 |         } catch (e) {
      13 |             console.error(e.stack)
    > 14 |             console.error(e.data)
         |                     ^
      15 |         }
      16 |     })
      17 |

      at Object.test (pub/utils/pay/rustore.test.js:14:21)
```_
