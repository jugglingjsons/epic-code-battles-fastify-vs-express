const fastify = require('fastify')({
  logger: process.env.LOGS
})

function sleep(ms) {
  return new Promise((resolve) => {
    setTimeout(resolve, ms);
  });
}
fastify.get('/', async (request, reply) => {
  // await sleep(process.env.TIMEOUT);
  return { hello: 'world' }
})

const start = async () => {
  try {
    await fastify.listen(process.env.PORT)
  } catch (err) {
    fastify.log.error(err)
    process.exit(1)
  }
}
start()