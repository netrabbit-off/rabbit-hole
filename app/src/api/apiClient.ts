class ApiError extends Error {
    constructor(message, status, body) {
        super(message);
        this.name = 'ApiError';
        this.status = status;
        this.body = body;
    }
}

class ApiClient {
    public baseUrl;
    constructor({ baseUrl } = {}) {
        if (!baseUrl) {
            throw new Error('ApiClient: baseUrl is required');
        }
        this.baseUrl = baseUrl.replace(/\/+$/, ''); // strip trailing slashes
    }

    _buildHeaders(extraHeaders = {}, skipAuth = false) {
        const headers = { Accept: 'application/json', ...extraHeaders };
        if (!skipAuth) {
            const token = this.getToken();
            if (token) {
                headers.Authorization = `Bearer ${token}`;
            }
        }
        return headers;
    }

    async _request(method, path, { body, headers, skipAuth, signal, raw } = {}) {
        const url = path.startsWith('http') ? path : `${this.baseUrl}${path}`;

        const finalHeaders = this._buildHeaders(headers, skipAuth);

        let payload = body;
        if (body !== undefined && !(body instanceof FormData)) {
            finalHeaders['Content-Type'] = finalHeaders['Content-Type'] ?? 'application/json';
            payload = typeof body === 'string' ? body : JSON.stringify(body);
        }

        let response;
        try {
            response = await fetch(url, {
                method,
                headers: finalHeaders,
                body: payload,
                signal,
            });
        } catch (e) {
            // Network error, DNS, offline, aborted request
            throw new ApiError(e.message || 'Network error', 0, null);
        }

        // Non-2xx responses
        if (!response.ok) {
            let errorBody = null;
            try {
                errorBody = await response.json();
            } catch {
                try {
                    errorBody = await response.text();
                } catch {
                    errorBody = null;
                }
            }
            const message = (errorBody && (errorBody.message || errorBody.error)) || `HTTP ${response.status}`;
            throw new ApiError(message, response.status, errorBody);
        }

        if (raw) return response;

        if (response.status === 204) return null;

        // Try JSON, fall back to text
        const contentType = response.headers.get('content-type') || '';
        if (contentType.includes('application/json')) {
            return response.json();
        }
        return response.text();
    }

    getToken = () => {
        return "123";
    }

    get(path, options) {
        return this._request('GET', path, options);
    }
    post(path, body, options = {}) {
        return this._request('POST', path, { ...options, body });
    }
    put(path, body, options = {}) {
        return this._request('PUT', path, { ...options, body });
    }
    patch(path, body, options = {}) {
        return this._request('PATCH', path, { ...options, body });
    }
    delete(path, options) {
        return this._request('DELETE', path, options);
    }

    url(path) {
        return path.startsWith('http') ? path : `${this.baseUrl}${path}`;
    }
    authHeaders() {
        const token = this.getToken();
        return token ? { Authorization: `Bearer ${token}` } : {};
    }
}

export { ApiClient, ApiError };
export default new ApiClient({baseUrl: "http://77.91.66.163:8088"});
