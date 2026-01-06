import axios from "axios";

const send = async ({ method = '', path = '', data = {}, token = '' } = {}) => {
  const baseUrl = 'http://localhost:8080/api'
  const url = `${baseUrl}${path}`

  const headers = {
    "Content-Type": "application/json",
    "Authorization": `Bearer ${token}`
  }

  const options = {
    method,
    url,
    headers,
    data
  }

  try {
    const response = await axios(options)
    return response.data
  } catch (error) {
    throw error
  }
}

export const getApi = ({ path = '', token = '' } = {}) => {
  return send({ method: 'GET', path, token })
}

export const putApi = ({ path = '', data = {}, token = '' } = {}) => {
  return send({ method: 'PUT', path, data, token })
}

export const postApi = ({ path = '', data = {}, token = '' } = {}) => {
  return send({ method: 'POST', path, data, token })
}

export const delApi = ({ path = '', data = {}, token = '' } = {}) => {
  return send({ method: 'DELETE', path, data, token })
}