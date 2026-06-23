FROM gcr.io/distroless/python3-debian13:nonroot

WORKDIR /app

COPY --chown=nonroot:nonroot main.py .python-version ./

ENV PORT=8080

EXPOSE 8080

CMD ["main.py"]
