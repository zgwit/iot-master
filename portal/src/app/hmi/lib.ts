export function SvgToDataUrl(svg: string) {
  return "data:image/svg+xml;base64," + atob(svg)
}

export function GetFieldDeeply(obj: any, key: string) {
  let keys = key.split('.')
  for (let i = 0; i < keys.length; i++) {
    let key = keys[i]
    if (!obj.hasOwnProperty(key))
      return undefined
    obj = obj[key]
  }
  return obj
}
