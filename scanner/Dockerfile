FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . ./

RUN go build -o scanner .

# ---------- Runtime Stage ----------
FROM debian:bookworm-slim AS runtime

# Install OS packages
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        ca-certificates \
        nmap \
        gobuster \
        ffuf \
        sqlmap \
        curl \
        git \
        perl \
        unzip \
        dnsutils \
        golang && \
    rm -rf /var/lib/apt/lists/*

# Install Nikto
RUN git clone https://github.com/sullo/nikto.git /opt/nikto && \
    ln -s /opt/nikto/program/nikto.pl /usr/local/bin/nikto && \
    chmod +x /opt/nikto/program/nikto.pl

# Build and install Amass (arch-safe)
RUN git clone https://github.com/owasp/amass.git /opt/amass && \
    cd /opt/amass && go build -o /usr/local/bin/amass ./cmd/amass && \
    rm -rf /opt/amass

# Download wordlist
RUN mkdir -p /usr/share/wordlists/dirb && \
    curl -sSL https://raw.githubusercontent.com/danielmiessler/SecLists/master/Discovery/Web-Content/common.txt \
    -o /usr/share/wordlists/dirb/common.txt

# Copy scanner binary
COPY --from=builder /app/scanner /usr/local/bin/scanner

ENTRYPOINT ["scanner"]
