FROM scratch
# Copy our static executable.
COPY dist/cer_linux_386/cer /go/bin/cer
# Run the hello binary.
ENTRYPOINT ["/go/bin/cer"]
