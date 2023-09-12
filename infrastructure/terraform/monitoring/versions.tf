terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "4.67.0"
    }

    elasticsearch = {
      source  = "phillbaker/elasticsearch"
      version = "2.0.7"
    }

    elasticstack = {
      source  = "elastic/elasticstack"
      version = "0.6.2"
    }

    gosoline = {
      source  = "justtrackio/gosoline"
      version = "1.2.1"
    }

    grafana = {
      source  = "grafana/grafana"
      version = "2.2.0"
    }
  }

  required_version = "1.5.5"
}
