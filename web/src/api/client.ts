export class ApiError extends Error {
  constructor(public status: number, public payload: unknown) {
    super(`HTTP ${status}`);
  }
}

async function handle<T>(res: Response): Promise<T> {
  if (!res.ok) {
    const payload = await res.json().catch(() => null);
    throw new ApiError(res.status, payload);
  }
  return res.json() as Promise<T>;
}

export const apiGet  = <T>(path: string) => fetch(path).then(handle<T>);
export const apiPost = <T>(path: string, body: unknown, token?: string) =>
  fetch(path, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json', ...(token ? { Authorization: token } : {}) },
    body: JSON.stringify(body),
  }).then(handle<T>);
export const apiPut = <T>(path: string, body: unknown, token: string) =>
  fetch(path, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json', Authorization: token },
    body: JSON.stringify(body),
  }).then(handle<T>);
export const apiDelete = <T>(path: string, token: string) =>
  fetch(path, {
    method: 'DELETE',
    headers: { Authorization: token },
  }).then(handle<T>);