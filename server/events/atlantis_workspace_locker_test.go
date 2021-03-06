package events_test

import (
	"testing"

	"github.com/runatlantis/atlantis/server/events"
	. "github.com/runatlantis/atlantis/testing"
)

var repo = "repo/owner"
var workspace = "default"

func TestTryLock(t *testing.T) {
	locker := events.NewDefaultAtlantisWorkspaceLocker()

	t.Log("the first lock should succeed")
	Equals(t, true, locker.TryLock(repo, workspace, 1))

	t.Log("now another lock for the same repo, workspace, and pull should fail")
	Equals(t, false, locker.TryLock(repo, workspace, 1))
}

func TestTryLockDifferentWorkspaces(t *testing.T) {
	locker := events.NewDefaultAtlantisWorkspaceLocker()

	t.Log("a lock for the same repo and pull but different workspace should succeed")
	Equals(t, true, locker.TryLock(repo, workspace, 1))
	Equals(t, true, locker.TryLock(repo, "new-workspace", 1))

	t.Log("and both should now be locked")
	Equals(t, false, locker.TryLock(repo, workspace, 1))
	Equals(t, false, locker.TryLock(repo, "new-workspace", 1))
}

func TestTryLockDifferentRepo(t *testing.T) {
	locker := events.NewDefaultAtlantisWorkspaceLocker()

	t.Log("a lock for a different repo but the same workspace and pull should succeed")
	Equals(t, true, locker.TryLock(repo, workspace, 1))
	newRepo := "owner/newrepo"
	Equals(t, true, locker.TryLock(newRepo, workspace, 1))

	t.Log("and both should now be locked")
	Equals(t, false, locker.TryLock(repo, workspace, 1))
	Equals(t, false, locker.TryLock(newRepo, workspace, 1))
}

func TestTryLockDifferent1(t *testing.T) {
	locker := events.NewDefaultAtlantisWorkspaceLocker()

	t.Log("a lock for a different pull but the same repo and workspace should succeed")
	Equals(t, true, locker.TryLock(repo, workspace, 1))
	new1 := 2
	Equals(t, true, locker.TryLock(repo, workspace, new1))

	t.Log("and both should now be locked")
	Equals(t, false, locker.TryLock(repo, workspace, 1))
	Equals(t, false, locker.TryLock(repo, workspace, new1))
}

func TestUnlock(t *testing.T) {
	locker := events.NewDefaultAtlantisWorkspaceLocker()

	t.Log("unlocking should work")
	Equals(t, true, locker.TryLock(repo, workspace, 1))
	locker.Unlock(repo, workspace, 1)
	Equals(t, true, locker.TryLock(repo, workspace, 1))
}

func TestUnlockDifferentWorkspaces(t *testing.T) {
	locker := events.NewDefaultAtlantisWorkspaceLocker()
	t.Log("unlocking should work for different workspaces")
	Equals(t, true, locker.TryLock(repo, workspace, 1))
	Equals(t, true, locker.TryLock(repo, "new-workspace", 1))
	locker.Unlock(repo, workspace, 1)
	locker.Unlock(repo, "new-workspace", 1)
	Equals(t, true, locker.TryLock(repo, workspace, 1))
	Equals(t, true, locker.TryLock(repo, "new-workspace", 1))
}

func TestUnlockDifferentRepos(t *testing.T) {
	locker := events.NewDefaultAtlantisWorkspaceLocker()
	t.Log("unlocking should work for different repos")
	Equals(t, true, locker.TryLock(repo, workspace, 1))
	newRepo := "owner/newrepo"
	Equals(t, true, locker.TryLock(newRepo, workspace, 1))
	locker.Unlock(repo, workspace, 1)
	locker.Unlock(newRepo, workspace, 1)
	Equals(t, true, locker.TryLock(repo, workspace, 1))
	Equals(t, true, locker.TryLock(newRepo, workspace, 1))
}

func TestUnlockDifferentPulls(t *testing.T) {
	locker := events.NewDefaultAtlantisWorkspaceLocker()
	t.Log("unlocking should work for different 1s")
	Equals(t, true, locker.TryLock(repo, workspace, 1))
	new1 := 2
	Equals(t, true, locker.TryLock(repo, workspace, new1))
	locker.Unlock(repo, workspace, 1)
	locker.Unlock(repo, workspace, new1)
	Equals(t, true, locker.TryLock(repo, workspace, 1))
	Equals(t, true, locker.TryLock(repo, workspace, new1))
}
