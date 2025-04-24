export function AddGuiones(texto: string): string {
  return texto.trim().replace(/\s+/g, "-");
}
