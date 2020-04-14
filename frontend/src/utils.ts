export async function FetchDataFromApi(url: string) {
    let api_host = process.env.VUE_APP_API_HOST != 'undefined' ? process.env.VUE_APP_API_HOST : '';
    let response = await fetch(api_host + url,
        {
            mode: 'cors', // no-cors, *cors, same-origin
            headers: {
                'Access-Control-Allow-Origin': '*'
            },
        }
    )
    return await response.json()
}
