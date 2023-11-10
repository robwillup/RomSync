package emulators

import "testing"

func TestFindPlatformFromExtension(t *testing.T) {
	expected := "snes"
	actual := FindEmulatorFromExtension("C:\\Users\\Gamer\\ReallyCoolGame.sfc")

	if expected != actual {
		t.Fatalf("Failed to FindPlatformFromExtension().\nexpected: %s\nactual: %s", expected, actual)
		return
	}

	expected = "megadrive"
	actual = FindEmulatorFromExtension("gamer/games/ReallyHardGame.md")

	if expected != actual {
		t.Fatalf("Failed to FindPlatformFromExtension().\nexpected: %s\nactual: %s", expected, actual)
		return
	}

	expected = "mastersystem"
	actual = FindEmulatorFromExtension("/home/gamer/ReallyNostalgicGame.sms")

	if expected != actual {
		t.Fatalf("Failed to FindPlatformFromExtension().\nexpected: %s\nactual: %s", expected, actual)
		return
	}

	expected = "gba"
	actual = FindEmulatorFromExtension("C:\\Users\\Gamer\\ReallyShortGame.gba")

	if expected != actual {
		t.Fatalf("Failed to FindPlatformFromExtension().\nexpected: %s\nactual: %s", expected, actual)
		return
	}

	expected = "atari2600"
	actual = FindEmulatorFromExtension("C:\\Users\\Gamer\\ReallyOld.a26")

	if expected != actual {
		t.Fatalf("Failed to FindPlatformFromExtension().\nexpected: %s\nactual: %s", expected, actual)
		return
	}

	expected = "sega32x"
	actual = FindEmulatorFromExtension("/home/gamer/ReallyRareGame.32x")

	if expected != actual {
		t.Fatalf("Failed to FindPlatformFromExtension().\nexpected: %s\nactual: %s", expected, actual)
		return
	}
}

func TestFindPlatformFromExtension_GameCube(t *testing.T) {
	expected := "gc"
	actual := FindEmulatorFromExtension("C:\\Users\\Gamer\\AGameCubeGame.gcm")

	if expected != actual {
		t.Fatalf("Failed to FindPlatformFromExtension_GameCube().\nexpected: %s\nactual: %s", expected, actual)
		return
	}

	actual = FindEmulatorFromExtension("/home/pi/AnotherGameCubeGame.gcz")

	if expected != actual {
		t.Fatalf("Failed to FindPlatformFromExtension_GameCube().\nexpected: %s\nactual: %s", expected, actual)
		return
	}

}