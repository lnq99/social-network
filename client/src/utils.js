export function squarePreview() {
  let imgs = document.getElementsByClassName('preview')
  for (let img of imgs) img.style.height = `${img.clientWidth}px`
}

export function squarePreviewDiv() {
  let imgs = document.querySelectorAll('.preview')
  for (let img of imgs) {
    img.style.maxHeight = `${img.clientWidth}px`
    img.style.minHeight = `${img.clientWidth}px`
  }
}

export function handlerResponse(response) {
  let responseOK =
    response && response.status === 200 && response.statusText === 'OK'
  if (responseOK) {
    let data = response.data
    console.log(data)
    if (data) {
      return data
    }
  }
}

export function getCookie(name) {
  const value = `; ${document.cookie}`
  const parts = value.split(`; ${name}=`)
  if (parts.length === 2)
    return parts
      .pop()
      .split(';')
      .shift()
}

export function overwriteCookie(name, val) {
  document.cookie = `${name}=${val}; path=/`
}
