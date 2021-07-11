
export const handleErrors = (response) => {
    if (!response.ok) {
        const errorMessage = "Failed to save session. Please try again.";
        console.log(`Error saving session: ${response.status} - ${response.url}`);
        throw Error(errorMessage);
    }
    return response;
};

const postData = (url, data) => {
    return fetch(url, {
        method: 'POST',
        mode: 'cors',
        credentials: 'same-origin',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    });
};

const saveSession = (userId, start, end, sessionName) => {
    const url = `/v1/users/${userId}/session`;

    const body = {
        start: start.toISOString(),
        end: end.toISOString(),
        name: sessionName
    };

    return postData(url, body)
};

export const client = { saveSession };
