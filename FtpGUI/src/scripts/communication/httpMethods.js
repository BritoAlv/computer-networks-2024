export async function get(url) {
    let response;

    await fetch(url)
    .then(data => {
        if(!data.ok) {
            response =  new Error("There was an error while fetching")
        }
        else {
            response = data.json()
        }
    });
    return response;
}

export async function post(url, request) {
    let response;

    await fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(request),
    })
    .then(data => {
        if(!data.ok) {
            response =  new Error("There was an error while fetching")
        }
        else {
            response = data.json()
        }
    });

    return response;
}