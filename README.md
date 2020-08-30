
# viacaocometa-alert

viacaocometa.com.br alerter to check ticket availability

# Motivation

I created this code to check the ticket availability on the weekends. Normally, because of pandemic (COVID-19) it is getting pretty hard to get more ticket hour options. So, instead of checking manually their site, I created this script thing to warm me when they will increase ticket hour options.


# How does it work

The simple as counting 1, 2, 3.

1. Will load configuration based on the .env file (I'll explain how this works further in this documentation)
2. Find next weekend (Sunday is hardcoded, but you may change)
3. Make search request to Viação Cometa site
4. Send result formatted to a Whatsapp number

# .env File

There are only 3 environment values that are really needed. Here is an example file:

``` sh
ORIGIN_ID=999
DESTINATION_ID=999
WHATSAPP_RECIPIENT=9999999999@s.whatsapp.net
```

## Origin and Destination code

You may get directly from their site or do some simple http GET request. E.g., below is the result (part of it) of all the destination cities starting from Sao Paulo (Brazil).


``` sh
$ http https://www.viacaocometa.com.br/content/jca/cometa/pt-br/jcr:content.getDestinos.json?origem=467
HTTP/1.1 200 OK
X-Vhost: www.viacaocometa.com.br

{
    "listaLocalidade": {
        "erro": {
            "descricaoAlerta": "",
            "descricaoErro": "",
            "ocorreuAlerta": false,
            "ocorreuErro": false,
            "sessionValid": true
        },
        "lsLocalidade": [
            {
                "cidade": "A PRATA",
                "id": 443,
                "uf": "SP"
            },
            {
                "cidade": "AMERICANA",
                "id": 64,
                "uf": "SP"
            },
            ...
             {
                "cidade": "ÁGUAS DA PRATA",
                "id": 14,
                "uf": "SP"
            }
        ]
    }
}           
```

## WhatsApp recipient

It's your direct number without any symbols. So, if your destination have a number like this: `1 (202) 358-0001 `

You will simply put: 

``` sh
WHATSAPP_RECIPIENT=12023580001@s.whatsapp.net
```
