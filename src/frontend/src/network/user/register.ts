export interface IRegister {
  name: string
  password: string
}

export async function register(data: IRegister): Promise<any> {
  const url = '/api/user/register'
  try {
    const res = await fetch(url, {
      method: 'POST',
      mode: 'cors',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    })
    return res.json()
  } catch (err) {
    console.error(err)
  }
}
