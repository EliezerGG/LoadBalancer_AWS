import express, { Request, Response } from 'express';

const app = express();
const PORT = 3000;

app.get('/check', (_req: Request, res: Response) => {
  res.sendStatus(200);
});

app.get('/', (_req: Request, res: Response) => {
  res.json({
    Instancia: 'Instancia #1 - API #1',
    Curso: 'Seminario de Sistemas 1',
    Estudiante: 'Eliezer Guevara - 202100179',
  });
});

app.get('/get-data', (_req: Request, res: Response) => {
  res.json({
    Instancia:  'Instancia #1 - API #1',
    Curso:      'Seminario de Sistemas 1',
    Estudiante: 'Eliezer Guevara - 202100179',
    Lenguaje:   'Node.js / TypeScript',
    Puerto:     PORT,
  });
});

app.get('/health', (_req: Request, res: Response) => {
  res.status(200).json({ status: 'ok', api: 'API-1' });
});

app.listen(PORT, () => {
  console.log(`API #1 (TypeScript) corriendo en puerto ${PORT}`);
});