FROM gcr.io/distroless/nodejs24-debian13:nonroot

WORKDIR /app

COPY --chown=nonroot:nonroot package.json index.js ./

ENV NODE_ENV=production
ENV PORT=8080

EXPOSE 8080

CMD ["index.js"]
