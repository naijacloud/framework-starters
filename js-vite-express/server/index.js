import express from 'express';
import { fileURLToPath } from 'node:url';
import path from 'node:path';
const dir = path.dirname(fileURLToPath(import.meta.url));
const app = express();
app.get('/api/health', (_q, r) => r.json({ ok: 'TGL10OK' }));
app.use(express.static(path.join(dir, '../dist')));
app.listen(process.env.PORT || 3000);
