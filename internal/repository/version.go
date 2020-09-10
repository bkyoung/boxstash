package repository

type VersionRepository interface {
    CreateVersion()
    ReadVersion()
    UpdateVersion()
    DeleteVersion()
    LatestVersion()
    ListVersions()
}
