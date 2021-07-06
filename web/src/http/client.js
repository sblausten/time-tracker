

const saveSession = () => {
    const url = `/v1/users/:userId/session`;

    return fetch(url)
        .then(res => {
            return res.json();
        });
};

export const client = { saveSession };
