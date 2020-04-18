export async function FetchDataFromApi(url: string, page: bigint, itemsPerPage: bigint) {
    let api_host = process.env.VUE_APP_API_HOST != 'undefined' ? process.env.VUE_APP_API_HOST : '';
    let url_for_fetch = new URL(api_host + url)
    url_for_fetch.search = new URLSearchParams({
        'page': page.toString(),
        'itemsPerPage': itemsPerPage.toString()
    }).toString()
    let response = await fetch(url_for_fetch.toString(),
        {
            mode: 'cors', // no-cors, *cors, same-origin
            headers: {
                'Access-Control-Allow-Origin': '*'
            },
        }
    )
    let total = response.headers.get("X-Total-Count");
    return [await response.json(), total == null ? 0 : parseInt(total)]
}
