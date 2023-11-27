import { RouteContext } from "$fresh/server.ts";

export interface Readings {
  ClientID: string;
  Topic: string;
  Payload: string;
  Timestamp: string;
}

export default async function Page(_req: Request, ctx: RouteContext) {
  const resp = await fetch(
    `http://localhost:4444/list`,
  );

  if (!resp.ok) {
    return <h1>An Error occurred</h1>;
  }

  const readings: Readings[] = await resp.json();

  return (
    <div>
      <h1>Readings</h1>
      <table>
        <thead>
          <tr>
            <th>ClientID</th>
            <th>Topic</th>
            <th>Payload</th>
            <th>Timestamp</th>
          </tr>
        </thead>
        <tbody>
          {readings.map((reading) => (
            <tr>
              <td>{reading.ClientID}</td>
              <td>{reading.Topic}</td>
              <td>{reading.Payload}</td>
              <td>{reading.Timestamp}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}
