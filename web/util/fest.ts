import { escape } from "querystring";

export function generateUniqueFestId(name: string): string {
  let id = name.toLocaleLowerCase();
  id = id.replaceAll(/[^a-zA-Z\d]/g, " ");
  id = id.replaceAll(" ", "").slice(0, 18);

  return id;
}
