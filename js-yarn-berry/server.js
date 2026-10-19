import express from 'express';
const app = express();
app.get('/', (_req, res) => res.send('<h1>TGL10OK yarn-berry</h1>'));
app.listen(process.env.PORT || 3000, () => console.log('listening'));
