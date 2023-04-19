# **Contributing**

## **Commiting code**

This project utilise [convetional commits](https://www.conventionalcommits.org/en/v1.0.0/):

- Major version bump prefix: `BREAKING CHANGE:`, `feat!:`, `refactor!:`
- Minor version bump prefix: `feat:`, `refactor:`
- Patch version bump prefix: `fix:`

NB: these rules only apply after 1.0.0. While < 1.0.0 the commit messages are "downshifted" one level.

Manual bump to 1.0.0:

```bash
git commit --allow-empty -m "chore: release 1.0.0" -m "release-as: 1.0.0"
```

## **Permissions**

In GitHub you must change the repository settings: `settings -> actions[general] -> Allow GitHub Actions to create and approve pull requests reviews`

## **Release Please**

[release-please](https://github.com/googleapis/release-please)
[release-please-action](https://github.com/google-github-actions/release-please-action)
[release-please-action-manifest](https://github.com/googleapis/release-please/blob/main/docs/manifest-releaser.md)
