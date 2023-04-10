# hte-danger-zone-job

### Setup
1. Build image
   ```
   docker build -t danger-zone-job .
   ```
2. Run image
   ```
   docker run -d --name danger-zone-job \
     -e REDIS_HOST='redis:6379' \
     -e REDIS_PASSWORD='' \
     -e REDIS_CHANNEL_GEOLOC='c-geoloc' \
     -e REDIS_CHANNEL_CREATE_ZONE='c-dangerzone:create' \
     --network hte \
     danger-zone-job
   ```

### Redis key convention

Streams: `s-<NAME>`
Stream Last ID read: `slid-<STREAM_NAME>:<SERVICE>`
Queues: `q-<NAME>`
Individual values: `v-<NAME>`
Channels: `c-<NAME>`

# Workflow
